# ASCII Art Web - Docker Team Plan

Split responsibilities for 3 team members to containerize the web application.

---

## Team Member 1: Dockerfile & Build

**Focus**: Create the Dockerfile with best practices

### Tasks
- [ ] Create multi-stage Dockerfile
  - Stage 1: Build stage (use `golang:1.22-alpine`)
  - Stage 2: Production stage (use `alpine:latest`)
- [ ] Copy only necessary files (main.go, go.mod, handlers/, Artprinter/, banners/, templates/)
- [ ] Expose port 8080
- [ ] Add LABEL metadata:
  - `maintainer="team"`
  - `version="1.0"`
  - `description="ASCII Art Web Generator"`
- [ ] Use non-root user for security
- [ ] Add `.dockerignore` file (exclude .git, tests, docs)

**Deliverable**: `Dockerfile` + `.dockerignore`

---

## Team Member 2: Image & Container Management

**Focus**: Build image, run container, apply best practices

### Tasks
- [ ] Build Docker image with proper tagging
  ```bash
  docker build -t ascii-art-web:1.0 .
  docker tag ascii-art-web:1.0 ascii-art-web:latest
  ```
- [ ] Run container with metadata labels
  ```bash
  docker run -d --name ascii-art-container \
    --label project=ascii-art \
    --label environment=development \
    -p 8080:8080 ascii-art-web:1.0
  ```
- [ ] Verify container is running and accessible
- [ ] Document image size and optimize if needed

**Deliverable**: Running container + build/run commands documented

---

## Team Member 3: Cleanup & Documentation

**Focus**: Garbage collection, README updates, verification

### Tasks
- [ ] Create cleanup script for unused objects
  ```bash
  # Remove stopped containers
  docker container prune -f
  # Remove unused images
  docker image prune -f
  # Remove dangling images
  docker rmi $(docker images -f "dangling=true" -q)
  ```
- [ ] Update README.md with Docker section:
  - Build instructions
  - Run instructions
  - Cleanup commands
- [ ] Verify all metadata labels are applied correctly
  ```bash
  docker inspect ascii-art-container --format='{{.Config.Labels}}'
  ```
- [ ] Test full workflow: build → run → test → cleanup

**Deliverable**: Cleanup script + updated README.md

---

## Verification Checklist (Team Review)
- [ ] Dockerfile follows best practices (multi-stage, minimal image)
- [ ] Image has proper tags and labels
- [ ] Container runs and serves on port 8080
- [ ] Cleanup commands work without errors
- [ ] README documents all Docker commands
