class Response {
  constructor(body) {
    this.body = body;
  }
  text() {
    let that = this;
    return that.body.toString();
  }
}

new Response("this body");
