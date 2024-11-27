# Radiofy 🎵

A real-time music sharing platform where users can create clubs, invite friends, and share their favorite music together. Radiofy integrates with Spotify to provide a seamless music sharing experience.

## Features ✨

- **Club Creation**: Create your own music club and invite friends
- **Real-time Music Sharing**: Share music in real-time with club members
- **Spotify Integration**: Seamlessly connect with Spotify for music playback
- **Chat System**: Communicate with club members while sharing music
- **User Authentication**: Secure user authentication system
- **Responsive UI**: Modern and user-friendly interface

## Tech Stack 🛠️

### Backend
- Go (Golang)
- MySQL 8.0
- WebSocket for real-time communication
- JWT Authentication
- Docker

### Frontend
- Vue.js 3
- Tailwind CSS
- Spotify Web API
- Axios for HTTP requests
- Vue Router
- Vuex for state management

## Prerequisites 📋

- Docker
- Go 1.16 or higher
- Node.js and npm
- Spotify Developer Account

## Installation & Setup 🚀

1. **Clone the repository**
   ```bash
   git clone https://github.com/batuhannoz/radiofy.git
   cd radiofy
   ```

2. **Set up the database**
   ```bash
   docker run --name radiofy-database \
   -e MYSQL_ROOT_USER=root \
   -e MYSQL_ROOT_PASSWORD=password \
   -e MYSQL_DATABASE=radiofy \
   -v $(pwd)/database/init.ddl:/docker-entrypoint-initdb.d/1-ddl.sql \
   -p 3307:3306 \
   -d mysql:8.0.23 \
   --default-authentication-plugin=mysql_native_password \
   --character-set-server=utf8mb4 \
   --collation-server=utf8mb4_general_ci
   ```

3. **Start the backend**
   ```bash
   cd backend
   APP_ENV=dev go run main.go
   ```

4. **Start the frontend**
   ```bash
   cd frontend
   npm install
   npm run serve
   ```

## Running the Application 🎮

1. Access the application at `http://localhost:8080`
2. Create an account or log in
3. Create a club or join an existing one
4. Start sharing music with your friends!

## Screenshots 📸

### Club Interface
![Club Interface](https://github.com/batuhannoz/radiofy/blob/main/img/Club.png)

### Live Demo
![Radiofy Demo](https://github.com/batuhannoz/radiofy/blob/main/img/Radiofy.gif)

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request.

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Troubleshooting 🔧

If you encounter Docker permission issues:
```bash
sudo chmod 666 /var/run/docker.sock
```

To clean up Docker containers:
```bash
docker rm -f $(docker ps -aq)
```
