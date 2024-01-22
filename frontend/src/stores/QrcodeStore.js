import { writable } from 'svelte/store';

export const qrcodeContent = writable({
  data: "",
  fgColor: "#ffffff",
  bgColor: "#000000",
});