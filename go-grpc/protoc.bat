bin\protoc --go_out=./api/v1/ --go_opt=paths=source_relative --go-grpc_out=./api/v1/ --go-grpc_opt=paths=source_relative  --proto_path=./api/v1/ ./api/v1/api.proto

for /f %%f in ('dir /b \api\v1*.proto' ) do (
    echo %%~nf
)