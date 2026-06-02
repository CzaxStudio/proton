using System;
using System.Diagnostics;
using System.Threading;

namespace proton
{
    class protlib
    {
        public static void Main(string[] args)
        {
            Console.WriteLine("please wait.......");
            Thread.Sleep(2000);
            Console.WriteLine("build for proton completed for windows");
            Thread.Sleep(1000);
            Console.WriteLine("this is a test file to see if the ide is working properly");
            Thread.Sleep(1500);
            Process.Start(@"C:\Users\DELL\AppData\Local\Programs\Microsoft VS Code\Code.exe");
            // This is just a test file for testing the IDE, not related to the Proton Framework
            // This is not a core library file
        }
    }
}
