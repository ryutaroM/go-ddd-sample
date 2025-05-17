# Domain-Driven Design (DDD) Sample Repository in Go

This repository serves as a sample implementation of Domain-Driven Design (DDD) principles in Go. It demonstrates how to organize a Go application following DDD architectural patterns.

## Overview

This project showcases a payment processing system built using DDD concepts. It provides a practical example of how to structure domain logic, application services, and infrastructure components while maintaining clear boundaries between different layers of the application.

## Project Structure

```
/
├── api/                                # Interface/Presentation Layer
│   └── http/                           # HTTP API implementation
│       ├── handler/                    # Request handlers
│       ├── request/                    # Request parsing and validation
│       └── response/                   # Response formatting
├── cmd/                                # Application entry points
│   └── main.go                         # Main application bootstrap
├── internal/                           # Private application code
│   ├── acl/                            # Anti-Corruption Layer (translation between contexts)
│   │   └── payment/user/               # User-Payment context translation
│   ├── domain/                         # Domain Layer (core business logic)
│   │   ├── payment/                    # Payment domain model
│   │   └── user/                       # User domain model
│   ├── infrastructure/                 # Infrastructure Layer
│   │   └── persistence/                # Repository implementations
│   └── usecase/                        # Application Layer (Use Cases)
│       └── payment/                    # Payment services
│           ├── command/                # Command objects (write operations)
│           └── query/                  # Query objects (read operations)
```

## Key DDD Concepts Demonstrated

- **Bounded Contexts**: Separate contexts for payment and user domains
- **Entities & Value Objects**: Domain models with proper validation and behavior
- **Repositories**: Abstraction for data access with domain-specific interfaces
- **Application Services**: Orchestrating domain objects to fulfill use cases
- **Anti-Corruption Layer (ACL)**: Translating between bounded contexts
- **Command Query Responsibility Segregation (CQRS)**: Separation of read and write operations

## Implementation Details

- **Domain Layer**: Contains the core business logic and domain models
  - Value objects (PaymentID, Amount, Quantity)
  - Entities with validation logic (Payment)
  - Repository interfaces

- **Application Layer**: Implements use cases using domain objects
  - Services that orchestrate domain operations
  - Commands and queries following CQRS pattern

- **Infrastructure Layer**: Provides implementations for technical concerns
  - Repository implementations for data persistence

- **Interface Layer**: Handles user interaction with the system
  - HTTP handlers for API endpoints
  - Request/response transformations

## How to Use This Repository

This repository serves as a reference implementation for Go developers wanting to apply DDD principles in their projects. Study the code organization, separation of concerns, and interactions between different layers to understand how DDD can be practically implemented in Go.

## Getting Started

1. Clone the repository
2. Run the application:
   ```
   go run cmd/main.go
   ```
3. Access the API endpoints:
   - GET `/payments/{id}` - Find a payment by ID
   - POST `/payments` - Create a new payment

## License

[MIT License](LICENSE)
