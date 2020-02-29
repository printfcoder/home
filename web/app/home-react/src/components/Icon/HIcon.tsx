import * as React from "react";
import PropTypes from 'prop-types';

import "../../assets/icon/iconfont.js"

const HIcon = ({name, color, size = 30}: { name: string, color?: string, size?: number }) => (
    <svg aria-hidden="true" style={{
        /*  fontFamily: "iconfont",
          fontSize: "16px",
          fontStyle: "normal",*/
        verticalAlign: "-0.15em",
        fill: "currentColor",
        overflow: "hidden"
    }} className="icon" fill={color} width={size} height={size}>
        <use xlinkHref={`#icon-${name}`}/>
    </svg>
);

HIcon.propTypes = {
    name: PropTypes.string.isRequired,
    color: PropTypes.string,
    size: PropTypes.number
};


export default HIcon;
