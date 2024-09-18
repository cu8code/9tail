import { Link } from "react-router-dom";

export default function Project() {
    return (
        <li>
            <Link to={"/editor"}>One</Link>
            <Link to={"/editor"}>Two</Link>
        </li>
    )
}