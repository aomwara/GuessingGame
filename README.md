# GuessingGame
Agoda/Intern/GuessingGame

Task:
- FE
    - Login if not authenticated
    - Guess the number
        - Do API call to backend to guess the number
- BE
    - API endpoints
        - `/login`
            - Very simple yes / no with password combination
            - Return "token"
        - `/guess`
            - Access to this endpoint needed to be authenticated via token returned from login
            - Guess the hidden number - if correct, return HTTP 201 and regenerate the number
    - RESTful
    - Your response should be in form of JSON format
    - Responses should have CRUD functionality
Bonus (for challenge):
- FE
    - Use React.js context for authentication
- BE
    - Use of middleware for authentication
    - If we wanted to hide the guess data by not using GET, can we use other method to do so
