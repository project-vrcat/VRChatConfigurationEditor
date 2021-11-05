import { postMessage, postMessageAsync } from "./webview2";

export function openURL(url: string) {
  postMessage("openURL", [url]);
}

export async function selectDirectory(title: string): Promise<string> {
  return postMessageAsync("selectDirectory", [title]);
}

export async function loadConfigFile(): Promise<string> {
  return postMessageAsync("loadConfigFile", []);
}

export async function saveConfigFile(data: string): Promise<boolean> {
  return new Promise<boolean>((resolve, reject) => {
    postMessageAsync("saveConfigFile", [data]).then((i) => {
      if (i === "success") return resolve(true);
      reject(i);
    }).catch((err) => {
      reject(err);
    });
  });
}
