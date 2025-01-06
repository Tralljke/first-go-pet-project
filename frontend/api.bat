java -jar ..\swagger-codegen-cli.jar generate ^
    -i http://localhost:8080/swagger/doc.json ^
    -l typescript-fetch ^
    -o src\api
echo API client generated successfully