import * as React from "react";
import {ReactComponent as Icon} from '../../assets/icon/iconfont.svg';
import PropTypes from 'prop-types';

const HIcon = ({name, color, size}: { name: string, color: string, size: number }) => (
    <svg className={`icon icon-${name}`} fill={color} width={size} height={size}>
        <use xlinkHref={`${Icon}#icon-${name}`}/>
    </svg>
);

HIcon.propTypes = {
    name: PropTypes.string.isRequired,
    color: PropTypes.string,
    size: PropTypes.number
};


export default HIcon;
