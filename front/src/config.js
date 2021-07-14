const isProductionEnv = process.env.NODE_ENV === "production";

export const API_DOMAIN = isProductionEnv
  ? "http://localhost:8000/api"
  : "http://localhost:8000";

export const BASE_URL = isProductionEnv ? "/project/" : "/";
