go / websockets / fiber / redis streams


cfg & env

`.env.example` хранит в себе пример для `.env`

нужно переименовать или скопировать в `.env`

дефолтный порт 8080

дефолтный jwt secret "a-string-secret-at-least-256-bits-long"


DOCKER

note: дефолтный порт 8080

сбилдить образы и запусить `docker compose up --build`

или

`
docker compose build app
`

`
docker compose up
`

или

попасть в контейнер
`docker compose run --rm app bash`


GO

точка входа `./cmd/app/main.go`

получить бинарник `go build -o server ./cmd/app/main.go`



TEST

в папке `/web` лежит простенький html+js для протыкивания руками

валидные прегены jwt токенов (вставить в форму), в токене зашит user_id

user_id=101 & 
token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAxIn0.2nbyHB2XmSbtk_UfFfcP3rXjOolCSdwGkO9rtuUEexg

user_id=102 & 
token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAyIn0.IijsmEuksoCXxpbflf_Kz4zgKJ3K2tNHb9qsIHZd210


если нажать на connect, когда соединение уже есть, то в форме останется ошибка (нельзя под одним токеном два коннекта), старый коннект потеряется (но будет жить, это видно по логам)

