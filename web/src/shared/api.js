import jwtInterceptor from "./jwt.interceptor";

export async function SetSubsystem(name, enable) {
    let payload = {
        subsystem: name,
        enable: enable
    }
    await jwtInterceptor.post("/api/set-subsystem", payload, {
        withCredentials: true,
        credentials: "include",
        headers: {
            'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
        },
    }).then(function (response) {
        // handle success
        //console.log(response);
        return response
    }).catch(function (error) {
        // handle error
        //console.log(error);
        return error
    })
}