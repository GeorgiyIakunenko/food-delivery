export const getImageUrl = (url) => {
  return new URL(`../assets/images/${url}`, import.meta.url).href;
};
