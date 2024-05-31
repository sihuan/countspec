import sys


# 400.perlbench 	C 	PERL Programming Language
# 401.bzip2 	C 	Compression
# 403.gcc 	C 	C Compiler
# 429.mcf 	C 	Combinatorial Optimization
# 445.gobmk 	C 	Artificial Intelligence: go
# 456.hmmer 	C 	Search Gene Sequence
# 458.sjeng 	C 	Artificial Intelligence: chess
# 462.libquantum 	C 	Physics: Quantum Computing
# 464.h264ref 	C 	Video Compression
# 471.omnetpp 	C++ 	Discrete Event Simulation
# 473.astar 	C++ 	Path-finding Algorithms
# 483.xalancbmk 	C++ 	XML Processing 

intcase = ["400", "401", "403", "429", "445", "456", "458", "462", "464", "471", "473", "483"]

# 410.bwaves 	Fortran 	Fluid Dynamics
# 416.gamess 	Fortran 	Quantum Chemistry
# 433.milc 	C 	Physics: Quantum Chromodynamics
# 434.zeusmp 	Fortran 	Physics / CFD
# 435.gromacs 	C/Fortran 	Biochemistry/Molecular Dynamics
# 436.cactusADM 	C/Fortran 	Physics / General Relativity
# 437.leslie3d 	Fortran 	Fluid Dynamics
# 444.namd 	C++ 	Biology / Molecular Dynamics
# 447.dealII 	C++ 	Finite Element Analysis
# 450.soplex 	C++ 	Linear Programming, Optimization
# 453.povray 	C++ 	Image Ray-tracing
# 454.calculix 	C/Fortran 	Structural Mechanics
# 459.GemsFDTD 	Fortran 	Computational Electromagnetics
# 465.tonto 	Fortran 	Quantum Chemistry
# 470.lbm 	C 	Fluid Dynamics
# 481.wrf 	C/Fortran 	Weather Prediction
# 482.sphinx3 	C 	Speech recognition 

fpcase = ["410", "416", "433", "434", "435", "436", "437", "444", "447", "450", "453", "454", "459", "465", "470", "481", "482"]

def read_file(file_path):
    intfile_path = f"int_{file_path}"
    fpfile_path = f"fp_{file_path}"
    try:
        intentries = []
        fpentries = []
        with open(file_path, 'r') as file:
            l = ''
            for line in file:
                if line.split(".")[0] in intcase:
                    intentries.append(line)
                elif line.split(".")[0] in fpcase:
                    fpentries.append(line)
        # 排序并写入 csv 文件
        intentries.sort()
        fpentries.sort()
        with open(intfile_path, 'w') as f:
            for entry in intentries:
                f.write(entry)

        with open(fpfile_path, 'w') as f:
            for entry in fpentries:
                f.write(entry)
                    
            
    except FileNotFoundError:
        print("File not found!")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <file_path>")
    else:
        file_path = sys.argv[1]
        read_file(file_path)
