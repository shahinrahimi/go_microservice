import React from "react";
import { Toast } from "react-bootstrap";
interface ToasterInterface {
    message: string
    show: boolean
}
const Toaster: React.FC<ToasterInterface> = ({ message, show }) => {
    const [visibility, setVisibility] = React.useState(false)

    React.useEffect(() => {
        setVisibility(show)
    }, [show])

    return (
        <Toast onClose={() => setVisibility(false)} show={visibility} delay={3000} autohide>
            <Toast.Header>
                <strong className="mr-auto">File Upload</strong>
            </Toast.Header>
            <Toast.Body>{message}</Toast.Body>
        </Toast>
    )
}

export default Toaster;