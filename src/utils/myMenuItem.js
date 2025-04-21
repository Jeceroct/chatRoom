class myMenuItem {
  title;
  submitFunc;
  funcArgs;

  constructor(title, submitFunc, ...funcArgs) {
    this.title = title;
    this.submitFunc = submitFunc;
    this.funcArgs = funcArgs;
  }

  submit() {
    this.submitFunc(...this.funcArgs);
  }

  getTitle() {
    return this.title;
  }

  setTitle(title) {
    this.title = title;
  }

  setSubmitFunc(submitFunc) {
    this.submitFunc = submitFunc;
  }

  setFuncArgs(...funcArgs) {
    this.funcArgs = funcArgs;
  }

  getSubmitFunc() {
    return this.submitFunc;
  }

  getFuncArgs() {
    return this.funcArgs;
  }

  getMenuItem() {
    return this;
  }

}

export default myMenuItem;