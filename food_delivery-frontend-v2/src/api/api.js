import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

export const getProducts = () => {
  return api
    .get("/products")
    .then((response) => {
      console.log(response.data);
      return response.data;
    })
    .catch((error) => {
      throw error;
    });
};

export const login = (email, password) => {
  return api
    .post("/auth/login", {
      email: email,
      password: password,
    })
    .then((response) => {
      console.log(response.data);
      return response.data;
    })
    .catch((error) => {
      throw error;
    });
};
