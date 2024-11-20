# Devbook Frontend - A Tailored Go-Based Social Network Project

This repository contains the frontend portion of the **Devbook** project, a simple social network application designed to explore the capabilities of the Go programming language in web development. 
Originally based on a Udemy course, I adapted and expanded the project to include my personal preferences and a more modern tech stack.

---

## About the Project

Devbook showcases key social network features such as user authentication, user and post management, and follower interactions. 
The frontend is built with Go templates and minimal JavaScript for responsiveness. Unlike the course's original implementation, 
which relied on Bootstrap and plain CSS, this version employs **Tailwind CSS** via the Tailwind CLI, offering a lightweight and customizable styling approach.

Additionally, the course content was in Portuguese, but this implementation is entirely in English, including all code, comments, and UI elements. 
This adaptation ensures broader accessibility while allowing me to practice and apply international development standards.

---

## Features

- **Authentication**: Secure login and registration with Go's backend handling.
- **Posts**: CRUD operations for posts, including illustrative "like" and "dislike" functionality.
- **Follow System**: Follow and unfollow users, with modals displaying followers and following lists.
- **Responsive Design**: Fully responsive UI powered by **Tailwind CSS**.
- **Minimal JavaScript**: Replaced Bootstrap JS from the course with custom lightweight solutions using jQuery.

---

## Tech Stack

### Frontend
- **Go**: Template rendering and server-side logic.
- **Tailwind CSS**: For responsive and modern styling.
- **jQuery**: Simplified JavaScript for interactivity.
- **HTML**: Templates creation and display.

### Backend Integration
This frontend integrates seamlessly with the Devbook backend (not included here), which features:
- RESTful API endpoints.
- MySQL database for data persistence.
- Secure user authentication.

---

### Prerequisites
- Install [Go](https://go.dev/doc/install) (v1.18+ recommended).
- Install [Tailwind CSS CLI](https://tailwindcss.com/docs/installation).
- Install [Air](https://github.com/cosmtrek/air) for live reloading during development.
- Set up the **Devbook Backend** and its database.
