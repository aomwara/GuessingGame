async function authChecker(token) {
    return await fetch('http://localhost:8080/authCheck', {
      method: 'GET',
      headers: {
        'Token': token
      },
    }).then(response => response.json())
    .then((responseData) => {
        return responseData.status
    })
}
export default authChecker