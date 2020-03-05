import * as React from "react";

export interface NotFoundProps {
}

class state {
}

export class NotFound extends React.Component<NotFoundProps, {}> {
    constructor(props: any) {
        super(props);
    }
    state: state = {};
    render() {
        return (
            <div>
                404
            </div>
        );
    }
}