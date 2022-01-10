import { Link } from "react-router-dom";

import classes from "./MainNavigation.module.css";

export default function MainNavigation() {
  return (
    <header className={classes.header}>
      <div className={classes.logo}>COVID-19 Admission Recommendation AI</div>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="result">Result</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
}
