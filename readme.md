# Introduction

This is an interview task about a game that mentions an OXO series game, which is similar to tic-tac-toe but might have different rules. The task involves designing some APIs for the game, including player management system, game room management system, endless challenge system, game log collector, and payment processing system. These APIs are based on RESTful style, with basic functionalities such as listing all players, adding a player, listing all game rooms, adding a game room, listing all reservations, and adding a reservation, etc. The APIs should also provide detailed documentation, including example requests and responses. Additionally, some logic needs to be implemented, like challenge logic in the endless challenge system and payment logic in the payment processing system. The APIs should provide both unit and integration tests to ensure complete and stable functionalities.

## Task

Use Golang to write a RESTful API for managing players and levels, game rooms and reservations, endless challenges, game logs, and payments in the OXO series game.

## API Development Challenge

Please use Golang or OpenResty to create an API server, and write a `Dockerfile` and `docker-compose.yml`, ensuring the project can start using `docker-compose up -d`. If a database or auxiliary application is needed, include it in the `docker-compose.yml`.

### 1. Player Management System

Your task is to design and implement a RESTful API to manage players and levels in the OXO series game.

**Requirements**:

1. Endpoints
   - `/players`:
     - `GET`: List all players, returning a JSON list containing each player's ID, name, and level info.
     - `POST`: Register a new player, accepting a JSON request containing the player's name and level. Returns the new player's ID.
   - `/players/{id}`:
     - `GET`: Get detailed info for a specific player ID.
     - `PUT`: Update the information of a specific player ID.
     - `DELETE`: Delete a specific player ID.
   - `/levels`:
     - `GET`: List all levels, returning a JSON list containing each level's ID and name.
     - `POST`: Add a new level, accepting a JSON request containing the level's name. Returns the new level's ID.

**Hints**:

1. Ensure API design follows RESTful principles with a clear structure.
2. Use appropriate HTTP status codes to indicate operation results.
3. Provide detailed API documentation, including example requests and responses.
4. Test your API to ensure complete and stable functionality.

### 2. Game Room Management System

Your task is to design and implement a RESTful API to manage game rooms and reservations in the OXO series game.

**Requirements**:

1. Endpoints
   - `/rooms`:
     - `GET`: List all game rooms, returning a JSON list containing each room's ID, name, and status info.
     - `POST`: Add a new game room, accepting a JSON request containing the room's name and description. Returns the new room's ID.
   - `/rooms/{id}`:
     - `GET`: Get detailed info for a specific room ID.
     - `PUT`: Update the information of a specific room ID.
     - `DELETE`: Delete a specific room ID.
   - `/reservations`:
     - `GET`: Query game room reservations with the following query parameters:
       - `room_id`: (optional) Specify the room ID to query reservations for, if not provided, query all rooms' reservations.
       - `date`: (optional) Specify the date to query.
       - `limit`: (optional) Specify the maximum number of reservations to return.
     - Return JSON formatted reservation results, containing each reservation's ID, room ID, date, time, and player info.
     - `POST`: Add a new game room reservation, accepting a JSON request containing room ID, date, time, and player info. Returns the new reservation's ID.

**Hints**:

1. Ensure API design is simple and easy to use, offering extensive filtering and querying capabilities.
2. Provide detailed API documentation, including example requests and responses.
3. Test your API to ensure complete and stable functionality.
4. Use unit tests and integration tests to verify your logic and API functionality.

### 3. Endless Challenge System

Your task is to design and implement a RESTful API to manage endless challenges in the OXO series game.

**Requirements**:

1. Endpoints

   - `/challenges`:
     - `POST`: Player joins a challenge, accepting a JSON request containing player ID and payment amount (fixed at 20.01). Returns the challenge participation status.
   - `/challenges/results`:
     - `GET`: List recent challenge results as a JSON list containing each challenge's ID, player ID, whether they won the jackpot, etc.

2. Logic
   - Each challenge lasts for 30 seconds, and a player can join every minute.
   - Each challenge requires a payment of 20.01.
   - After 30 seconds, there's a 1% chance for the challenger to win the entire jackpot.
   - Challenges can continue endlessly, with more participation increasing the player's chances of winning.

**Hints**:

1. Ensure API is simple and can handle high-frequency challenge requests.
2. Provide detailed API documentation, including example requests and responses.
3. Test your API to ensure complete and stable functionality.
4. Use unit tests and integration tests to verify logic and API functionality.

### 4. Game Log Collector

Your task is to design and implement a RESTful API to record each operation of players in the OXO series game.

**Requirements**:

1. Endpoints
   - `/logs`:
     - `GET`: Query game logs with the following query parameters:
       - `player_id`: (optional) Specify the player ID to query logs for, if not provided, query all players' logs.
       - `action`: (optional) Specify the action type to query.
         - Register
         - Login
         - Logout
         - Enter Room
         - Exit Room
         - Participate in Challenge
         - Challenge Result
       - `start_time`, `end_time`: (optional) Specify the time range for the query.
       - `limit`: (optional) Specify the maximum number of log entries to return.
     - Return JSON formatted log entries containing each operation's ID, player ID, action type, timestamp, and details.
     - `POST`: Add a new game operation log, accepting a JSON request containing player ID, action type, and details. Returns the new log ID.

**Hints**:

1. Ensure API design is simple and easy to use, with extensive filtering and querying capabilities.
2. Provide detailed API documentation, including example requests and responses.
3. Test your API to ensure complete and stable functionality.
4. Use unit tests and integration tests to verify logic and API functionality.

### 5. Payment Processing System

Your task is to design and implement a RESTful API to process various payment methods in the OXO series game, including credit card, bank transfer, third-party payment, and blockchain payment.

**Requirements**:

1. Endpoints

   - `/payments`:
     - `POST`: Process payment, accepting a JSON request containing payment method (credit card, bank transfer, third-party payment, blockchain payment), payment amount, and payment details. Returns payment results, including payment status and transaction ID.
   - `/payments/{id}`:
     - `GET`: Get detailed info for a specific payment.

2. Payment Processing Logic
   - Depending on the payment method, call different payment services:
     - Credit card payment: Simulate call to credit card payment gateway.
     - Bank transfer: Simulate call to bank transfer service.
     - Third-party payment: Simulate call to third-party payment platform.
     - Blockchain payment: Simulate call to blockchain payment gateway.
   - After successful payment, return a transaction ID.
   - In case of payment failure, return error information.

**Hints**:

1. Ensure API design is simple and easy, with high security.
2. Provide detailed API documentation, including example requests and responses.
3. Test your API to ensure complete and stable functionality.
4. Use unit tests and integration tests to verify payment logic and API functionality.
