# Team Member 1: Docker Guide

An interactive, step-by-step guide to create your Dockerfile. Complete each step, understand why, then check it off.

---

## Step 1: Create .dockerignore

**What**: A file telling Docker which files to SKIP when building.

**Why**: Smaller image, faster builds, no unnecessary files (tests, git history).

### Action
Create `C:\ascii-art\.dockerignore` with this content:

```
.git
*.md
*_test.go
TESTING_TODO.md
DOCKER_TEAM_PLAN.md
```

### Verify
```bash
# File should exist
dir .dockerignore
```

- [ ] Done - I created .dockerignore

---

## Step 2: Start the Dockerfile (Build Stage)

**What**: Multi-stage build - first stage compiles Go code.

**Why**: 
- Go needs ~1GB image to compile
- Final app only needs ~15MB runtime
- Multi-stage = compile in big image, copy binary to tiny image

### Action
Create `C:\ascii-art\Dockerfile` with:

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency file first (cache optimization)
COPY go.mod ./

# Copy source code
COPY main.go ./
COPY Artprinter/ ./Artprinter/
COPY handlers/ ./handlers/

# Build the binary
RUN go build -o ascii-art-web .
```

### Understanding Each Line

| Line | Purpose |
|------|---------|
| `FROM golang:1.22-alpine AS builder` | Use Go image, name it "builder" |
| `WORKDIR /app` | All commands run from /app directory |
| `COPY go.mod ./` | Copy dependencies first (Docker caches this layer) |
| `COPY ... ./` | Copy source code into container |
| `RUN go build -o ascii-art-web .` | Compile Go code into single binary |

- [ ] Done - I created the build stage

---

## Step 3: Add Production Stage

**What**: Second stage - tiny image with just our binary.

**Why**: 
- Alpine Linux is only ~5MB
- We only need the compiled binary + static files
- Smaller = faster deploys, less attack surface

### Action
Add this BELOW your existing Dockerfile content:

```dockerfile
# Stage 2: Production
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/ascii-art-web .

# Copy static files (templates and banners)
COPY templates/ ./templates/
COPY banners/ ./banners/

# Expose port
EXPOSE 8080

# Run the application
CMD ["./ascii-art-web"]
```

### Understanding Each Line

| Line | Purpose |
|------|---------|
| `FROM alpine:latest` | Fresh, tiny Linux image (~5MB) |
| `COPY --from=builder` | Copy binary FROM the first stage |
| `COPY templates/` | Copy HTML templates |
| `COPY banners/` | Copy ASCII art banner files |
| `EXPOSE 8080` | Document which port the app uses |
| `CMD [...]` | Command to run when container starts |

- [ ] Done - I added the production stage

---

## Step 4: Add Metadata Labels

**What**: Labels = key-value metadata attached to the image.

**Why**: Required by task. Helps identify images, track versions, know maintainers.

### Action
Add these lines AFTER `FROM alpine:latest`:

```dockerfile
LABEL maintainer="Dixon Osure"
LABEL version="1.0"
LABEL description="ASCII Art Web Generator"
LABEL project="ascii-art-web"
```

- [ ] Done - I added labels

---

## Step 5: Add Security (Non-root User)

**What**: Run app as regular user, not root.

**Why**: If attacker compromises app, they can't access system as root.

### Action
Add these lines BEFORE the `CMD` line:

```dockerfile
# Create non-root user
RUN adduser -D appuser
USER appuser
```

- [ ] Done - I added non-root user

---

## Final Dockerfile

Your complete `Dockerfile` should look like this:

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./
COPY Artprinter/ ./Artprinter/
COPY handlers/ ./handlers/

RUN go build -o ascii-art-web .

# Stage 2: Production
FROM alpine:latest

LABEL maintainer="Dixon Osure"
LABEL version="1.0"
LABEL description="ASCII Art Web Generator"
LABEL project="ascii-art-web"

WORKDIR /app

COPY --from=builder /app/ascii-art-web .
COPY templates/ ./templates/
COPY banners/ ./banners/

RUN adduser -D appuser
USER appuser

EXPOSE 8080

CMD ["./ascii-art-web"]
```

---

## Step 6: Test Your Build

**What**: Build the image to verify Dockerfile works.

**Why**: Catch errors before passing to Member 2.

### Action
Run in terminal:
```bash
cd C:\ascii-art
docker build -t ascii-art-web:test .
```

### Expected Output
```
 => [builder 1/6] FROM golang:1.22-alpine
 => [builder 2/6] WORKDIR /app
 ...
 => [stage-1 6/7] RUN adduser -D appuser
 => exporting to image
Successfully built ...
Successfully tagged ascii-art-web:test
```

- [ ] Done - Build succeeded!

---

## Checklist Summary

- [ ] `.dockerignore` created
- [ ] `Dockerfile` created with multi-stage build
- [ ] Labels added (maintainer, version, description, project)
- [ ] Non-root user configured
- [ ] Build test passed

**Hand off to Member 2** when all boxes are checked!
