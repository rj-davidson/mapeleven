import React from "react";
import './Button.css';
import { Link } from "react-router-dom";

const styles = ['button--primary', 'button--outline'];
const sizes = ['button--small', 'button--medium', 'button--large'];

export const Button = ({ children, type, onClick, buttonStyle, buttonSize }) => {
  const checkButtonStyle = styles.includes(buttonStyle) ? buttonStyle : styles[0];

  const checkButtonSize = sizes.includes(buttonSize) ? buttonSize : sizes[0];

  return (
    <Link to='/graph' className='button-graph'>
      <button
        className={`btn ${checkButtonStyle} ${checkButtonSize}`}
        onClick={onClick}
        type={type}
      >
        {children}
      </button>
    </Link>
  );
}