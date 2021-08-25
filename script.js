class Response {
  constructor(body) {
    this.body = body;
  }
  text() {
    return this.body.toString();
  }
  // text(body) {
  //   return body.toString();
  // }
}

new Response("this body");
