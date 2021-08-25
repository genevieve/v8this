function text() {
  return "the plain old function text";
}
class Response {
  constructor(body) {
    this.body = body;
  }
  text() {
    return this.body.toString();
  }
}

new Response("Hello, World");
