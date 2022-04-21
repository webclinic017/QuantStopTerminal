import axios from "axios";
const jwtInterceptor = axios.create({});


jwtInterceptor.interceptors.request.use((config) => {
  return config;
});

jwtInterceptor.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    if (error.response.status === 401) {
     var response = await axios.get(
         "/api/refresh-token",
        {
          withCredentials: true,
        }
      )
      .catch((err) => {
        return Promise.reject(err);
      });
      if(response && response.data){
        return axios(error.config);

      }
      else{
        return Promise.reject(error);
      }
    } else {
      return Promise.reject(error);
    }
  }
);

export default jwtInterceptor;
