/**
 * @return {string}
 * */
export function GetDatetimeStamp({hour12 = false}) {
  return new Date().toLocaleString(undefined, {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    hour12: hour12,
    minute: '2-digit',
    second: '2-digit'
  })
}
