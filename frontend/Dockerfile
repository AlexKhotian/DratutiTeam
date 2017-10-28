FROM node:boron  
RUN mkdir -p /app  
RUN mkdir -p /app/src
RUN mkdir -p /app/node_modules  
WORKDIR /app  
COPY package.json .  
COPY index.html . 
COPY webpack.config.js . 
RUN npm install   
COPY src ./src  
EXPOSE 8080  
CMD [ "npm", "run", "dev"]  