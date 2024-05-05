import sys


# 500.perlbench_r#1
# instructions: 1214375604958
# stdout: checkspam.2500.5.25.11.150.1.1.1.1.out
# stderr: checkspam.2500.5.25.11.150.1.1.1.1.err
# 500.perlbench_r#2
# instructions: 746949057779
# stdout: diffmail.4.800.10.17.19.300.out
# stderr: diffmail.4.800.10.17.19.300.err
# 500.perlbench_r#3

def read_file(file_path):
    try:
        entries = []
        with open(file_path, 'r') as file:
            l = ''
            for line in file:
                if line.startswith("stdout:") or line.startswith("stderr:"):
                    continue
                if line.startswith("instructions:") or line.startswith("Status: failed"):
                    ins = line.split(":")[1].strip() if line.startswith("instructions:") else 0
                    l = l + f',{ins}'
                    entries.append(l)
                    l = ''
                else:
                    l = line.strip()
        # 排序并写入 csv 文件
        entries.sort()
        with open('mydata.csv', 'w') as f:
            for entry in entries:
                f.write(entry + '\n')

        print(entries)
                    
            
    except FileNotFoundError:
        print("File not found!")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <file_path>")
    else:
        file_path = sys.argv[1]
        read_file(file_path)
