async function authChecker(token) {
    return await fetch('http://localhost:8080/authCheck', {
      method: 'GET',
      headers: {
        'Token': token
      },
      credentials: 'same-origin',
    }).then(response => response.json())
    .then((responseData) => {
        //return responseData
        if(responseData.status == "true"){
            return 0
        }else{
            return 1
        }
    })
}
export default authChecker