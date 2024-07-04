import { Sarabun } from "next/font/google";
import "./globals.css";

const sarabun = Sarabun({
  subsets: ['thai'], 
  variable: '--sarabun-font', 
  display: 'swap',
  weight: '400',
  style:'normal',
});

export const metadata = {
  title: "Thai-Wrong-Word Game",
  description: "quiz app for misspelled words in Thai language",
  icons: {
    icon: [
      { url: '/favicon.png', type: 'image/png' },
    ],
  },
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className={sarabun.className}>
      <body>{children}</body>
    </html>
  );
}
