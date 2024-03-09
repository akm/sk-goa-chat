FROM node:20.9.0-bookworm-slim

# Install firebse-tools and emulators
RUN npm install -g firebase-tools@12.9.0

RUN mkdir -p /srv/firebase-emulators
WORKDIR /srv/firebase-emulators

COPY ./firebase.json ./firebase.json
COPY ./.firebaserc ./.firebaserc
