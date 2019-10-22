docker build -t toury_image .

docker run  -p 80:80 -p 40000:40000 -t toury_image