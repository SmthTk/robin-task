# Robin Task API

This project implements a task management API using the Go programming language and the Gin framework.

## Features

- **Authentication**: JWT-based authentication
- **Roles**: Admin and User roles with access control
- **Tasks**: Create, Read, Update, Delete (CRUD) tasks
- **Comments**: Users can comment on tasks, and only the creator of the comment can modify or delete it
- **Archiving**: Archive tasks without deleting them
- **Change Logs**: Track changes to tasks (e.g., name, description, status)
- **Rate Limiting**: API rate-limited to avoid abuse
- **CORS**: Cross-origin resource sharing is enabled

## Getting Started

### Prerequisites

- Docker
- Go (for local development)
- MySQL (Docker or a local setup)

### Running with Docker

1. Clone the repository
   ```bash
   git clone https://github.com/your-repo/robin-task.git
   cd robin-task
