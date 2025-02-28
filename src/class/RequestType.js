export default class RequestType {
  type = {
    'text': 'text',
    'image': 'image',
    'file': 'file'
  }
  user = {
    'id': '',
    'name': '',
    'level': 0,
    'avatar': '',
    'title': '',
    'titleColor': '',
    'phone': ''
  }
  quote = {
    'from': this.user,
    'type': '',
    'context': '',
    'time': ''
  }
  fileContext = {
    'Title': '',
    'Context': ''
  }
  requestType = {
    'type': '',
    'context': '',
    'time': '',
    'from': this.user,
    'quote': this.quote
  }

  constructor (type, context, time, from, quote) {
    this.setRequestType(type, context, time, from, quote)
  }

  getResult () {
    return this.requestType
  }

  setUser (user) {
    this.requestType.from = user
  }

  setQuote (requestType) {
    this.requestType.quote.context = requestType.context
    this.requestType.quote.type = requestType.type
    this.requestType.quote.from = requestType.from
    this.requestType.quote.time = requestType.time
  }

  setType (type) {
    this.requestType.type = type
  }

  setContext (context) {
    this.requestType.context = context
  }

  setTime (time) {
    this.requestType.time = time
  }

  setRequestType (type, context, time, from, quote) {
    this.requestType.type = type
    this.requestType.context = context
    this.requestType.time = time
    this.requestType.from = from
    if (quote) {
      this.requestType.quote.context = quote.context
      this.requestType.quote.type = quote.type
      this.requestType.quote.from = quote.from
      this.requestType.quote.time = quote.time
    }
  }

  static User() {
    return new RequestType().user
  }

  static Quote() {
    return new RequestType().quote
  }

  static Type() {
    return new RequestType().type
  }

  static FileContext(name, context) {
    const fileContext = new RequestType().fileContext
    fileContext.Title = name
    fileContext.Context = context
    const str = JSON.stringify(fileContext)
    return str
  }
}