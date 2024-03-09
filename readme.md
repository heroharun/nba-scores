# NBA Real-Time Scores and Statistics Web Application

This web application simulates NBA matchDB in real-time and provides scores and statistics for each match.

## Installation and Configuration

Follow these steps to install and run the application:

1. Clone the repository:
    
    ```bash
   git clone https://github.com/your_username/nba-scores.git
    ```

2. Install Golang:
- Follow instructions from the [official Golang website](https://golang.org/doc/install) to install Golang on your system.

3. Run the application:
- Run the following command to start the application:

    ```bash
  cd nba-scores/backend
    go run main.go
    ```


4. Access the application:
- Please use postman in your local computer and use `http://localhost:8080` to open websocket connection.

## Usage

- View real-time scores and statistics of NBA matchDB.
- No need to refresh the page to get updated scores and statistics.

## API Endpoints

- `/ws`: Get updates
- `/`: Get details of sample usage.

## Project Structure

- `/backend`: Backend code written in Golang.
- `/frontend`: Frontend code written in HTML and CSS. Maybe added in the future.

## Technologies Used

- Backend: Golang
- Frontend: HTML, CSS 

## Contributors

- Harun Büyüktepe <harun.buyuktepe@gmail.com>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
