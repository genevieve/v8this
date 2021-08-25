class Response {
  constructor(body) {
    this.body = body;
  }
  text() {
    return this.body.toString();
  }
}

new Response("Hello, World");
