export function cleanImageUrl(url: string, width: number) {
  const cleanedUrl = url.replace(/(\?|\&)(q|auto|fit|ixlib|ixid)=[^&]+/g, "");
  const replacedUrl = cleanedUrl.replace(/(\?|\&)w=[^&]+/g, "");
  const modifiedUrl =
    replacedUrl + (replacedUrl.includes("?") ? "&" : "?") + "w=" + width;
  return modifiedUrl;
}
