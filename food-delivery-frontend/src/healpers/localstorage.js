export function setLocalStorageItem(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
}

export function getLocalStorageItem(key) {
    const item = localStorage.getItem(key);
    try {
        return item ? JSON.parse(item) : null;
    } catch (error) {
        console.error(`No JSON from localStorage for key "${key}":`);
        return null;
    }
}

export function removeLocalStorageItem(key) {
    localStorage.removeItem(key);
}