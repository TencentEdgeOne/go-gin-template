# Go Cloud Functions on EdgeOne Pages - Gin Framework

A full-stack demo application built with Next.js + Tailwind CSS frontend and Go Gin backend, showcasing how to deploy Go Cloud Functions using the Gin framework on EdgeOne Pages with RESTful API routing.

## 🚀 Features

- **Gin Framework Integration**: Full-featured Go web framework with middleware, JSON binding, and route groups
- **Modern UI Design**: Dark theme with #1c66e5 accent color, responsive layout with interactive elements
- **Interactive API Testing**: Built-in API endpoint panel — click "Call" to test each REST endpoint in real-time
- **RESTful API Design**: Complete CRUD operations with structured route groups (`/v1/users`, `/v1/posts`)
- **TypeScript Support**: Complete type definitions and type safety on the frontend

## 🛠️ Tech Stack

### Frontend
- **Next.js 15** - React full-stack framework (with Turbopack)
- **React 19** - User interface library
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS 4** - Utility-first CSS framework

### UI Components
- **shadcn/ui** - High-quality React components
- **Lucide React** - Beautiful icon library
- **class-variance-authority** - Component style variant management
- **clsx & tailwind-merge** - CSS class name merging utilities

### Backend
- **Go 1.21** - Cloud Functions runtime
- **Gin v1.10** - High-performance HTTP web framework for Go

## 📁 Project Structure

```
go-gin-template/
├── cloud-functions/                # Go Cloud Functions source
│   ├── api.go                     # Gin app with all REST API routes
│   ├── go.mod                     # Go module definition
│   └── go.sum                     # Go dependency checksums
├── src/
│   ├── app/                       # Next.js App Router
│   │   ├── globals.css           # Global styles (dark theme)
│   │   ├── layout.tsx            # Root layout
│   │   └── page.tsx              # Main page (API testing UI)
│   ├── components/               # React components
│   │   └── ui/                   # UI base components
│   │       ├── button.tsx        # Button component
│   │       └── card.tsx          # Card component
│   └── lib/                      # Utility functions
│       └── utils.ts              # Common utilities (cn helper)
├── public/                        # Static assets
│   ├── eo-logo-blue.svg          # EdgeOne logo (blue)
│   └── eo-logo-white.svg         # EdgeOne logo (white)
├── package.json                   # Project configuration
└── README.md                     # Project documentation
```

## 🚀 Quick Start

### Requirements

- Node.js 18+
- pnpm (recommended) or npm
- Go 1.21+ (for local development)

### Install Dependencies

```bash
pnpm install
# or
npm install
```

### Development Mode

```bash
edgeone pages dev
```

Visit [http://localhost:8088](http://localhost:8088) to view the application.

### Build Production Version

```bash
edgeone pages build
```

## 🎯 Core Features

### 1. Gin REST API Routes

All API endpoints are defined in a single `cloud-functions/api.go` file using Gin's route groups:

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/api/v1/hello` | Welcome message |
| GET | `/api/v1/health` | Health check (includes Go runtime version) |
| GET | `/api/v1/users` | List all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| POST | `/api/v1/users` | Create new user (JSON body binding) |
| GET | `/api/v1/posts` | List all posts |
| GET | `/api/v1/posts/:id` | Get post by ID |

### 2. Interactive API Testing Panel

- 7 pre-configured API endpoint cards with "Call" buttons
- Real-time JSON response display with syntax highlighting
- POST request support with pre-filled JSON body
- Loading states and error handling

### 3. Gin Framework Convention

The Go backend uses Gin's standard patterns — route groups, JSON binding, and middleware:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    v1 := r.Group("/v1")
    {
        v1.GET("/hello", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "Hello from Go Gin!",
            })
        })

        v1.GET("/users", getUsersHandler)
        v1.GET("/users/:id", getUserByIdHandler)
        v1.POST("/users", createUserHandler)
    }

    r.Run(":9000")
}
```

### 4. Data Models

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Post struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
}
```

## 🔧 Configuration

### Tailwind CSS Configuration
The project uses Tailwind CSS 4 with custom color variables:

```css
:root {
  --primary: #1c66e5;        /* Primary color */
  --background: #000000;     /* Background color */
  --foreground: #ffffff;     /* Foreground color */
}
```

### Component Styling
Uses `class-variance-authority` to manage component style variants with multiple preset styles.

## 📚 Documentation

- **EdgeOne Pages Official Docs**: [https://pages.edgeone.ai/document/go-functions](https://pages.edgeone.ai/document/go-functions)
- **Gin Framework**: [https://gin-gonic.com/docs](https://gin-gonic.com/docs)
- **Next.js Documentation**: [https://nextjs.org/docs](https://nextjs.org/docs)
- **Tailwind CSS Documentation**: [https://tailwindcss.com/docs](https://tailwindcss.com/docs)

## 🚀 Deployment Guide

### EdgeOne Pages Deployment

1. Push code to GitHub repository
2. Create new project in EdgeOne Pages console
3. Select GitHub repository as source
4. Configure build command: `edgeone pages build`
5. Deploy project

### Go Gin Cloud Function

Create `cloud-functions/api.go` in project root with your Gin application:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/v1/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello from Go Gin on EdgeOne Pages!",
        })
    })

    r.Run(":9000")
}
```

## Deploy

[![Deploy with EdgeOne Pages](https://cdnstatic.tencentcs.com/edgeone/pages/deploy.svg)](https://edgeone.ai/pages/new?from=github&template=go-gin-template)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/github/choosealicense.com/blob/gh-pages/_licenses/mit.txt) file for details.
