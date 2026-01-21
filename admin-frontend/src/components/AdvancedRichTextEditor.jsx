import React, { useState } from "react";
import { EditorContent, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Image from "@tiptap/extension-image";
import TextStyle from "@tiptap/extension-text-style";
import Color from "@tiptap/extension-color";
import Underline from "@tiptap/extension-underline";
import Link from "@tiptap/extension-link";
import Table from "@tiptap/extension-table";
import TableRow from "@tiptap/extension-table-row";
import TableCell from "@tiptap/extension-table-cell";
import TableHeader from "@tiptap/extension-table-header";
import Placeholder from "@tiptap/extension-placeholder";
import CharacterCount from "@tiptap/extension-character-count";
import FontFamily from "@tiptap/extension-font-family";
import TextAlign from "@tiptap/extension-text-align";
import Highlight from "@tiptap/extension-highlight";

const AdvancedRichTextEditor = ({ value, onChange, maxLength = 10000, placeholder = "Write a detailed description..." }) => {
  const [showWordCount, setShowWordCount] = useState(true);
  const [fullscreen, setFullscreen] = useState(false);

  const editor = useEditor({
    extensions: [
      StarterKit,
      Underline,
      TextStyle,
      Color,
      Link.configure({
        openOnClick: true,
        autolink: true,
        linkOnPaste: true,
      }),
      Image.configure({ 
        inline: true,
        allowBase64: true, 
      }),
      Table.configure({
        resizable: true,
      }),
      TableRow,
      TableCell,
      TableHeader,
      Placeholder.configure({
        placeholder,
      }),
      CharacterCount.configure({
        limit: maxLength,
      }),
      FontFamily,
      TextAlign.configure({
        types: ['heading', 'paragraph'],
      }),
      Highlight.configure({
        multicolor: true,
      }),
    ],
    content: value || "",
    onUpdate: ({ editor, transaction }) => {
      if (!transaction.docChanged) return;
      onChange(editor.getHTML());
      // Auto-resize logic could be added here if needed, but we'll handle it via CSS
    },
    editorProps: {
      handlePaste(view, event) {
        event.preventDefault();

        const clipboardData = event.clipboardData || window.clipboardData;
        const htmlContent = clipboardData.getData("text/html");

        // Handle image paste directly from clipboard
        const items = (event.clipboardData || event.originalEvent.clipboardData).items;
        for (const item of items) {
          if (item.type.indexOf("image") === 0) {
            const blob = item.getAsFile();
            const reader = new FileReader();
            reader.onload = (e) => {
              editor.chain().focus().setImage({ src: e.target.result }).run();
            };
            reader.readAsDataURL(blob);
            return true;
          }
        }

        // Preserve HTML styles when pasting
        if (htmlContent) {
          editor.chain().focus().insertContent(htmlContent).run();
        } else {
          const text = clipboardData.getData("text/plain");
          editor.chain().focus().insertContent(text).run();
        }
      },
    },
  });

  if (!editor) return null;

  const handleClick = (callback) => (event) => {
    event.preventDefault();
    callback();
  };

  // Get editor stats
  const wordCount = editor.storage.characterCount?.words() || 0;
  const characterCount = editor.storage.characterCount?.characters() || 0;
  const percentage = Math.min(100, Math.ceil((characterCount / maxLength) * 100)) || 0;

  // Insert table
  const insertTable = () => {
    editor.chain().focus().insertTable({ rows: 3, cols: 3, withHeaderRow: true }).run();
  };

  // Clear editor content
  const clearContent = () => {
    if (confirm("Are you sure you want to clear all content?")) {
      editor.chain().focus().clearContent().run();
    }
  };

  // Undo/Redo
  const undo = () => editor.chain().focus().undo().run();
  const redo = () => editor.chain().focus().redo().run();

  return (
    <div className={`border rounded ${fullscreen ? 'fixed inset-0 z-50 bg-white p-4 overflow-auto' : ''}`}>
      {/* Main Toolbar */}
      <div className="mb-2 p-1 bg-gray-50 border-b flex flex-wrap gap-1">
        <div className="flex space-x-1 items-center mr-2">
          <button
            onClick={handleClick(() => editor.chain().focus().toggleBold().run())}
            className={`p-1 rounded ${editor.isActive("bold") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Bold"
          >
            <strong>B</strong>
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleItalic().run())}
            className={`p-1 rounded ${editor.isActive("italic") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Italic"
          >
            <em>I</em>
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleUnderline().run())}
            className={`p-1 rounded ${editor.isActive("underline") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Underline"
          >
            <u>U</u>
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleStrike().run())}
            className={`p-1 rounded ${editor.isActive("strike") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Strikethrough"
          >
            <s>S</s>
          </button>
        </div>

        <div className="flex space-x-1 items-center mr-2">
          <select
            onChange={(e) => editor.chain().focus().setFontFamily(e.target.value).run()}
            className="border rounded px-1 text-sm"
            title="Font Family"
          >
            <option value="Arial">Arial</option>
            <option value="Times New Roman">Times</option>
            <option value="Courier New">Courier</option>
            <option value="Georgia">Georgia</option>
            <option value="Verdana">Verdana</option>
          </select>
          
          <select
            onChange={(e) => {
              editor.chain().focus()
                .setMark('textStyle', { style: `font-size: ${e.target.value}` })
                .run();
            }}
            className="border rounded px-1 text-sm"
            title="Font Size"
          >
            <option value="12px">12px</option>
            <option value="14px">14px</option>
            <option value="16px">16px</option>
            <option value="18px">18px</option>
            <option value="20px">20px</option>
            <option value="24px">24px</option>
            <option value="32px">32px</option>
          </select>
          
          <input
            type="color"
            onChange={(e) => editor.chain().focus().setColor(e.target.value).run()}
            className="w-8 h-6 border rounded"
            title="Text Color"
          />
        </div>

        <div className="flex space-x-1 items-center mr-2">
          <button
            onClick={handleClick(() => editor.chain().focus().setTextAlign('left').run())}
            className={`p-1 rounded ${editor.isActive({ textAlign: 'left' }) ? "bg-gray-300" : "bg-gray-100"}`}
            title="Align Left"
          >
            ⬅️
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().setTextAlign('center').run())}
            className={`p-1 rounded ${editor.isActive({ textAlign: 'center' }) ? "bg-gray-300" : "bg-gray-100"}`}
            title="Align Center"
          >
            ⬆️
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().setTextAlign('right').run())}
            className={`p-1 rounded ${editor.isActive({ textAlign: 'right' }) ? "bg-gray-300" : "bg-gray-100"}`}
            title="Align Right"
          >
            ➡️
          </button>
        </div>

        <div className="flex space-x-1 items-center mr-2">
          <button
            onClick={handleClick(() => editor.chain().focus().toggleBulletList().run())}
            className={`p-1 rounded ${editor.isActive("bulletList") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Bullet List"
          >
            •
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleOrderedList().run())}
            className={`p-1 rounded ${editor.isActive("orderedList") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Numbered List"
          >
            1.
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleBlockquote().run())}
            className={`p-1 rounded ${editor.isActive("blockquote") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Quote"
          >
            ❝
          </button>
        </div>

        <div className="flex space-x-1 items-center mr-2">
          <button
            onClick={handleClick(() => {
              const url = prompt("Enter link URL:");
              if (url) {
                editor.chain().focus().extendMarkRange("link").setLink({ href: url }).run();
              }
            })}
            className={`p-1 rounded ${editor.isActive("link") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Add Link"
          >
            🔗
          </button>
          <button
            onClick={handleClick(() => {
              const url = prompt("Enter image URL:");
              if (url) editor.chain().focus().setImage({ src: url }).run();
            })}
            className="p-1 rounded bg-gray-100"
            title="Add Image"
          >
            🖼️
          </button>
          <button
            onClick={handleClick(insertTable)}
            className={`p-1 rounded ${editor.isActive("table") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Insert Table"
          >
            🏢
          </button>
          <button
            onClick={handleClick(() => editor.chain().focus().toggleHighlight().run())}
            className={`p-1 rounded ${editor.isActive("highlight") ? "bg-gray-300" : "bg-gray-100"}`}
            title="Highlight"
          >
            🖌️
          </button>
        </div>
      </div>

      {/* Secondary Toolbar */}
      <div className="flex justify-between px-2 py-1 bg-gray-50 border-b text-sm">
        <div className="flex gap-2">
          <button
            onClick={handleClick(undo)}
            disabled={!editor.can().undo()}
            className={`p-1 rounded ${!editor.can().undo() ? "opacity-50" : "bg-gray-100"}`}
            title="Undo"
          >
            ↩️
          </button>
          <button
            onClick={handleClick(redo)}
            disabled={!editor.can().redo()}
            className={`p-1 rounded ${!editor.can().redo() ? "opacity-50" : "bg-gray-100"}`}
            title="Redo"
          >
            ↪️
          </button>
          <button
            onClick={handleClick(clearContent)}
            className="p-1 rounded bg-gray-100"
            title="Clear All"
          >
            🗑️
          </button>
        </div>
        <div>
          <button
            onClick={() => setFullscreen(!fullscreen)}
            className="p-1 rounded bg-gray-100"
            title={fullscreen ? "Exit Fullscreen" : "Fullscreen Mode"}
          >
            {fullscreen ? '⬇️' : '⬆️'}
          </button>
        </div>
      </div>

      {/* Editor Content with Auto-Resize */}
      <div className="relative">
        <EditorContent 
          editor={editor} 
          className="border p-2 text-lg resize-none overflow-auto"
          style={{ 
            minHeight: fullscreen ? '80vh' : '100px', // Minimum height
            height: 'auto', // Allow height to adjust
          }}
        />
        <style jsx>{`
          .ProseMirror {
            min-height: 100px;
            height: auto;
            overflow: auto;
            resize: vertical; /* Optional: allows manual vertical resize */
          }
          .ProseMirror p.is-editor-empty:first-child::before {
            content: attr(data-placeholder);
            float: left;
            color: #adb5bd;
            pointer-events: none;
            height: 0;
          }
        `}</style>
      </div>

      {/* Word Count */}
      <div className="flex justify-between text-xs text-gray-500 mt-1 px-2">
        <div>
          <button 
            onClick={() => setShowWordCount(!showWordCount)} 
            className="text-blue-500 underline"
          >
            {showWordCount ? "Hide Stats" : "Show Stats"}
          </button>
        </div>
        {showWordCount && (
          <div className="flex items-center gap-3">
            <span>{wordCount} words</span>
            <span>{characterCount} characters</span>
            <div className="flex items-center">
              <div className="w-24 h-2 bg-gray-200 rounded-full overflow-hidden">
                <div 
                  className={`h-full ${percentage > 90 ? 'bg-red-500' : 'bg-green-500'}`} 
                  style={{ width: `${percentage}%` }} 
                />
              </div>
              <span className="ml-1">{percentage}%</span>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default AdvancedRichTextEditor;