import React, { useState, useRef } from 'react';
import { Button } from 'primereact/button';
import { PostMethod } from '../api/http';
import MonacoEditor from 'react-monaco-editor';
import { Dialog } from 'primereact/dialog';

const interpreterAPI = process.env.REACT_APP_API_URL_INTERPRETER;

const Home = () => {
    const [codeText, setCodeText] = useState('');
    const [consoleText, setConsoleText] = useState('');
    const [html, setHtml] = useState('');
    const [showHtml, setShowHtml] = useState(false);
    const [htmlErr, setHtmlErr] = useState('');
    const [showHtmlErr, setShowHtmlErr] = useState(false);
    const [svgContent, setSvgContent] = useState('');
    const [showSvgModal, setShowSvgModal] = useState(false);
    const [showHtmlErrModal, setShowHtmlErrModal] = useState(false);
    const [showHtmlModal, setShowHtmlModal] = useState(false);

    const uploadInputRef = useRef(null);

    const CompileInterpreter = async () => {
        const resp = await PostMethod(interpreterAPI + 'Interpreter', { Content: codeText });
        const svgContent = resp?.CstImage;
        await setSvgContent(svgContent);
        const newHtml = resp?.Tabla;
        await setHtml(newHtml);
        const newHtmlErr = resp?.TablaErr;
        await setHtmlErr(newHtmlErr);
        await setConsoleText(resp?.Output);
    };

    const handleFileUpload = () => {
        uploadInputRef.current.click();
    };

    const handleFileInputChange = e => {
        const file = e.target.files[0];
        const reader = new FileReader();

        reader.onload = (event) => {
            setCodeText(event.target.result);
        };

        if (file) {
            reader.readAsText(file);
        }
    };

    const toggleSvgDisplay = () => {
        setShowSvgModal(!showSvgModal);
        setShowHtml(false);
        setShowHtmlErr(false);
    };

    const toggleHtmlDisplay = () => {
        setShowHtmlModal(!showHtmlModal);
        setShowHtmlErr(false);
        setShowSvgModal(false);
    };
    
    const toggleHtmlErrDisplay = () => {
        setShowHtmlErrModal(!showHtmlErrModal);
        setShowHtml(false);
        setShowSvgModal(false);
    };

    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
            <h1>T-Swift</h1>
            
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%' }}>
                <div>
                    <Button label="UPLOAD" className="p-button-success" onClick={handleFileUpload} />
                    <input ref={uploadInputRef} type="file" accept=".swift" onChange={handleFileInputChange} style={{ display: 'none' }} />
                </div>
                <div>
                    <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                        <Button label="CST" className="p-button-success" style={{ marginRight: '1%', whiteSpace: 'nowrap' }} onClick={toggleSvgDisplay}/>
                        <Button label="ERRORES" className="p-button-success" style={{ marginRight: '1%', whiteSpace: 'nowrap' }} onClick={toggleHtmlErrDisplay} />
                        <Button label="TABLA DE SIMBOLOS" className="p-button-success" style={{ whiteSpace: 'nowrap' }} onClick={toggleHtmlDisplay} />
                    </div>
                </div>
            </div>
            <div style={{ display: 'flex', justifyContent: 'left', alignItems: 'left', width: '100%', paddingTop: '1%' }}>
                <div style={{ width: '40%', height: '60%', marginRight: '10%' }}>
                    <label>Input Code</label>
                    <MonacoEditor
                        height={'570px'}
                        width={'800px'}
                        language="swift"
                        theme="vs-dark"
                        value={codeText}
                        onChange={setCodeText}
                        options={{
                            fontFamily: "'Fira Code', monospace",
                            fontSize: 14,
                            wordWrap: "on",
                            colorDecorators: true,
                            lineNumbers: "on",
                            automaticLayout: true,
                            renderLineHighlight: "all",
                            selectionHighlight: true,
                            minimap: {
                                enabled: true
                            },
                            scrollbar: {
                                verticalScrollbarSize: 10,
                                horizontalScrollbarSize: 10
                            },
                            tokenColorCustomizations: {
                                variables: "#FF5733",
                                functions: "#00BFFF",
                                keywords: "#FFD700"
                            }
                        }}
                    />

                </div>
        
                <div style={{ width: '40%', height: '60%', marginLeft: '10%' }}>
                    <label>Console Output</label>
                    <MonacoEditor
                        height={'570px'}
                        width={'600px'}
                        language="swift"
                        theme="vs-dark"
                        value={consoleText}
                        options={{ readOnly: true, scrollBeyondLastLine: false }}

                    />
                </div>
            </div>
            <div style={{ display: 'flex', justifyContent: 'left', alignItems: 'left', width: '100%', marginLeft: '0%', paddingTop: '1%' }}>
                <Button label="RUN >>>" className="p-button-success" onClick={CompileInterpreter} style={{ marginBottom: '1%' }} />
            </div>
            {showHtml && <div dangerouslySetInnerHTML={{ __html: html }} />}
            {showHtmlErr && <div dangerouslySetInnerHTML={{ __html: htmlErr }} />}
            
            {/* Modal for CST Visualization */}
            <Dialog
                visible={showSvgModal}
                onHide={toggleSvgDisplay}
                style={{ width: '80vw', maxWidth: '1200px' }}
                header="CST Visualization"
                modal
                closable
                closeOnEscape
            >
                <div style={{ width: '100%', height: '500px', overflow: 'auto' }}>
                    <div
                        dangerouslySetInnerHTML={{ __html: svgContent }}
                        style={{ width: '100%', height: '100%', transition: 'width 0.5s, height 0.5s' }}
                    />
                </div>
            </Dialog>
            
            {/* Modal for Error Table */}
            <Dialog
                visible={showHtmlErrModal}
                onHide={toggleHtmlErrDisplay}
                style={{ width: '80vw', maxWidth: '1200px' }}
                header="Tabla de Errores"
                modal
                closable
                closeOnEscape
            >
                <div style={{ width: '100%', height: '500px', overflow: 'auto' }}>
                    <div dangerouslySetInnerHTML={{ __html: htmlErr }} />
                </div>
            </Dialog>
            
            {/* Modal for Symbol Table */}
            <Dialog
                visible={showHtmlModal}
                onHide={toggleHtmlDisplay}
                style={{ width: '80vw', maxWidth: '1200px' }}
                header="Tabla de Símbolos"
                modal
                closable
                closeOnEscape
            >
                <div style={{ width: '100%', height: '500px', overflow: 'auto' }}>
                    <div dangerouslySetInnerHTML={{ __html: html }} />
                </div>
            </Dialog>
        </div>
    );
};

export default Home;
