class Response {
  constructor(body) {
    this.body = body;
  }
  text() {
    return this.body.toString();
  }
}

new Response("this body");
