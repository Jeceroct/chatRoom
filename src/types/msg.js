import RequestType from "@/class/RequestType"

const msg = new RequestType(
  RequestType.Type().text,
  '',
  new Date().toLocaleTimeString(),
  RequestType.User(),
  null
)

export default msg