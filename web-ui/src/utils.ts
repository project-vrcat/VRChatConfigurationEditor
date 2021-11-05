export interface resolution {
  height: number;
  width: number;
}

export function name2resolution(name: string): resolution {
  switch (name) {
    case "720p":
      return { width: 1280, height: 720 };
    case "2k":
      return { width: 2560, height: 1440 };
    case "4k":
      return { width: 3840, height: 2160 };
    default:
      return { width: 1920, height: 1080 };
  }
}

export function resolution2name(r: resolution): string {
  if (r.height === 720 && r.width === 1280) return "720p";
  if (r.height === 1440 && r.width === 2560) return "2k";
  if (r.height === 2160 && r.width === 3840) return "4k";
  return "1080p";
}
