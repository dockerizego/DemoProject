FROM golang

RUN git clone https://github.com/dockerizego/DemoProject.git \
   && go get -u github.com/gorilla/mux \
   && cd DemoProject \
   && go build main.go \
   && mv main ..\
   && cd .. \
   && rm -rf DemoProject/ \
   && rm -rf github.com/ \
   && mkdir logs
   
CMD ["./main"]
