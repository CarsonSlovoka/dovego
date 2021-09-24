import {GetDatetimeStamp} from "../pkg/datetime.js"

(() => {
  window.onload = () => {
    console.log(GetDatetimeStamp({}))
    console.log(GetDatetimeStamp({hour12: true}))
  }
})()
