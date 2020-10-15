FROM mongo

COPY init.json /init.json
CMD mongoimport --host mongo-db --db mydog-api --collection codeImage --drop --type json --file /init.json --jsonArray