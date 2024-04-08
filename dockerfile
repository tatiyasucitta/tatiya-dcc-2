#base imagenya
FROM golang:1.22.0

RUN mkdir /app
ADD . /app
WORKDIR /app

#buat bangun aplikasi go di container
#bakal jalanin perintah go build yang bikin app go di dlm container yang hasilnya bakal disimpen dgn nama "bin" di current dir
RUN go build -o bin .
CMD ["/app/bin"]