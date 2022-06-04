# How to creat eJwt token :
1. go to [jwt.io](https://jwt.io/)
2. enter below details in decode section:
```Json
    header :
    {
        "alg": "HS256",
        "typ": "JWT"
    }

    Payload:

    {
        "secret": "supersecretkey"
    }

    Verify Signature Key :

    "supersecretkey"
```

3. once signature is verified, copy token from Encode section. token string like `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWNyZXQiOiJzdXBlcnNlY3JldGtleSJ9.q-VEQa_-EPfbyLHNnn18KYBoS8meWoGPIzIQzsfy5Z4`

***

# Api Request and Response :

1. merchant list request & response:
    ```Json
    Request :
        curl --location --request GET 'http://localhost:8081/api/v1/merchant/list?page=1' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWNyZXQiOiJzdXBlcnNlY3JldGtleSJ9.q-VEQa_-EPfbyLHNnn18KYBoS8meWoGPIzIQzsfy5Z4'

    Response:
    {
        "success": true,
        "responseCode": 200,
        "data": {
            "merchants": [
                {
                    "id": 6,
                    "name": "merchant 5",
                    "code": "merchant5",
                    "status": 1,
                    "created_at": "2022-06-04T09:07:44Z",
                    "updated_at": "2022-06-04T09:07:44Z"
                },
                {
                    "id": 5,
                    "name": "merchant 4",
                    "code": "merchant4",
                    "status": 1,
                    "created_at": "2022-06-04T00:37:39Z",
                    "updated_at": "2022-06-04T00:37:39Z"
                }
            ],
            "pagination": {
                "next_page": true,
                "previous_page": false,
                "page": 1
            }
        },
        "error": null
        }
    ```

2. merchant create request & response:s
    ```Json
    Request :
        curl --location --request POST 'http://localhost:8081/api/v1/merchant/create' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWNyZXQiOiJzdXBlcnNlY3JldGtleSJ9.q-VEQa_-EPfbyLHNnn18KYBoS8meWoGPIzIQzsfy5Z4' \
        --data-raw '{
            "name": "merchant 6",
            "code": "merchant6",
            "status": 1
        }'

    Response:
    {
        "success": true,
        "responseCode": 200,
        "data": "Merchant created successfully.",
        "error": null
    }
    ```

3. merchant team list resuest & response:
    ```Json
        Request :
            curl --location --request GET 'http://localhost:8081/api/v1/merchant-team/list?page=2&merchant_id=2' \
            --header 'Content-Type: application/json' \
            --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWNyZXQiOiJzdXBlcnNlY3JldGtleSJ9.q-VEQa_-EPfbyLHNnn18KYBoS8meWoGPIzIQzsfy5Z4'

        Response:
        {
            "success": true,
            "responseCode": 200,
            "data": {
                "merchant_teams": [
                    {
                        "id": 5,
                        "merchant_id": 2,
                        "name": "team2",
                        "email": "merchant2.team2@gmail.com",
                        "status": 1,
                        "created_at": "2022-06-04T02:02:12Z",
                        "updated_at": "2022-06-04T02:02:12Z"
                    },
                    {
                        "id": 4,
                        "merchant_id": 2,
                        "name": "team1",
                        "email": "merchant2.team1@gmail.com",
                        "status": 1,
                        "created_at": "2022-06-04T02:01:58Z",
                        "updated_at": "2022-06-04T02:01:58Z"
                    }
                ],
                "pagination": {
                    "next_page": true,
                    "previous_page": true,
                    "page": 2
                }
            },
            "error": null
        }
    ```

4. merchant team create request & response:
    ```Json
    Request :
        curl --location --request POST 'http://localhost:8081/api/v1/merchant-team/create' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZWNyZXQiOiJzdXBlcnNlY3JldGtleSJ9.q-VEQa_-EPfbyLHNnn18KYBoS8meWoGPIzIQzsfy5Z4' \
        --data-raw '{
            "merchant_id":2,
            "name": "team5",
            "email": "merchant2.team5@gmail.com",
            "status": 1
        }'

    Response:
    {
        "success": true,
        "responseCode": 200,
        "data": "Merchant team created successfully.",
        "error": null
    }
    ```