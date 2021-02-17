import App from '../App'
import  { Redirect } from 'react-router-dom'

export default function logout(){
    localStorage.clear()
    window.location.reload(false);
}