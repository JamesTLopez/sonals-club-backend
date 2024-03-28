# Specifies a parent image
FROM golang:1.22.1-alpine3.18 

# RUN mkdir /app
  
WORKDIR /app
# Copies everything from your root directory into /app
COPY . /app
 
# Installs Go dependencies
# RUN go mod download


 
# Builds your app with optional configuration
RUN  go build -o /sonals-club-backend
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD [ "/sonals-club-backend" ]