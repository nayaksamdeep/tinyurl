#set -x
#curl -X POST -H "Content-Type: application/json" -d @test.json http://localhost:8080/v1/Contact
#curl -X GET -H "Content-Type: application/json" http://localhost:7000/v1/ListURL
##curl -X POST -H "Content-Type: application/json" -d '{"id":"10", "url":"www.facebook.com"}' http://localhost:8080/v1/ConvertURL?url="google"

count=1;
while [ $count -le 50000 ]
do
curl -X POST -H "Content-Type: application/json" -d '{"url":"www.facebook.com", "tinyurl":"http://localhost:8080/v1/google.com"}' http://localhost:7000/v1/ConvertURL

#echo $count
((count++))
done
