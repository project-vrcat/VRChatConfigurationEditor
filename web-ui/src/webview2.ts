//@ts-nocheck
let values = {};

window.chrome.webview.addEventListener("message", (event: any) => {
  // console.log(event);
  if (!event.data || !event.data.Method) return;
  if (event.data.Args.length === 2) {
    values[event.data.Args[0]] = event.data.Args[1];
  }
  // console.log(values);
});

export async function postMessageAsync(
  method: string,
  args: string[],
): Promise<string> {
  return new Promise<string>((resolve, reject) => {
    let valueInterval;
    const key = method + Math.floor(Date.now() / 1000);
    const promiseTimeout = setTimeout(() => {
      clearInterval(valueInterval);
      reject("timeout");
    }, 30000);

    postMessage(method, [key].concat(args));

    valueInterval = setInterval(() => {
      if (values[key]) {
        clearTimeout(promiseTimeout);
        resolve(values[key]);
        delete (values[key]);
        clearInterval(valueInterval);
      }
    }, 500);
  });
}

export function postMessage(method: string, args: string[]) {
  window.chrome.webview.postMessage({ method: method, args: args });
}
