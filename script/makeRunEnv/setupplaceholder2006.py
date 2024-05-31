#!/bin/python
import os

# 500.perlbench_r 502.gcc_r 503.bwaves_r 505.mcf_r 507.cactuBSSN_r 508.namd_r 510.parest_r 511.povray_r 519.lbm_r 520.omnetpp_r 521.wrf_r 523.xalancbmk_r 525.x264_r 526.blender_r 527.cam4_r 531.deepsjeng_r 538.imagick_r 541.leela_r 544.nab_r 548.exchange2_r 549.fotonik3d_r 554.roms_r 557.xz_r 600.perlbench_s 602.gcc_s 603.bwaves_s 605.mcf_s 607.cactuBSSN_s 619.lbm_s 620.omnetpp_s 621.wrf_s 623.xalancbmk_s 625.x264_s 627.cam4_s 628.pop2_s 631.deepsjeng_s 638.imagick_s 641.leela_s 644.nab_s 648.exchange2_s 649.fotonik3d_s 654.roms_s 657.xz_s
# 5xx - rate
ratedir = ["500.perlbench_r", "502.gcc_r", "503.bwaves_r", "505.mcf_r", "507.cactuBSSN_r", "508.namd_r", "510.parest_r", "511.povray_r", "519.lbm_r", "520.omnetpp_r", "521.wrf_r", "523.xalancbmk_r", "525.x264_r", "526.blender_r", "527.cam4_r", "531.deepsjeng_r", "538.imagick_r", "541.leela_r", "544.nab_r", "548.exchange2_r", "549.fotonik3d_r", "554.roms_r", "557.xz_r"]
# 6xx - speed
speeddir = ["600.perlbench_s", "602.gcc_s", "603.bwaves_s", "605.mcf_s", "607.cactuBSSN_s", "619.lbm_s", "620.omnetpp_s", "621.wrf_s", "623.xalancbmk_s", "625.x264_s", "627.cam4_s", "628.pop2_s", "631.deepsjeng_s", "638.imagick_s", "641.leela_s", "644.nab_s", "648.exchange2_s", "649.fotonik3d_s", "654.roms_s", "657.xz_s"]

spec2006dir = ["400.perlbench", "401.bzip2", "403.gcc", "410.bwaves", "416.gamess", "429.mcf", "433.milc", "434.zeusmp", "435.gromacs", "436.cactusADM", "437.leslie3d", "444.namd", "445.gobmk", "447.dealII", "450.soplex", "453.povray", "454.calculix", "456.hmmer", "458.sjeng", "459.GemsFDTD", "462.libquantum", "464.h264ref", "465.tonto", "470.lbm", "471.omnetpp", "473.astar", "481.wrf", "482.sphinx3", "483.xalancbmk"]




def refdir(label, name):
    if name in ratedir:
        return f"run_base_refrate_{label}.0000"
    elif name in speeddir:
        return f"run_base_refspeed_{label}.0000"
    elif name in spec2006dir:
        return f"run_base_ref_{label}.0000"
    else:
        raise Exception(f"Unknown benchmark {name}")
    
def testdir(label, name):
    if name in ratedir + speeddir:
        return f"run_base_test_{label}.0000"
    elif name in spec2006dir:
        return f"run_base_test_{label}.0000"
    else:
        raise Exception(f"Unknown benchmark {name}")

# for name in ratedir + speeddir:
label2017 = "llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64"
label2006 = "gcc_s"
label = label2006
    
for name in spec2006dir:
    ref = refdir(label, name)
    print(f"Ref dir: {ref}")
    # 寻找 ref 文件夹里所有 名为 *.{label} 的可执行文件，删除并留下一个名为 *.mygo 的占位文件
    # 例如 500.perlbench_r/run_base_refrate_llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64.0000/perlbench_r_base.llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64 删除后，新建一个 perlbench_r_base.mygo 占位文件
    exe_files = [f for f in os.listdir(f"{name}/{ref}") if f.endswith(label)]
    for exe in exe_files:
        print(f"Remove {name}/{ref}/{exe}")
        os.remove(f"{name}/{ref}/{exe}")
        print(f"Create {name}/{ref}/{exe.replace("_base." + label, ".mygo")}")
        with open(f"{name}/{ref}/{exe.replace("_base." + label, ".mygo")}", "w") as f:
            pass
    # 编辑 ref 文件夹下的 mygocmd 文件，把 其中的 ref 文件夹名改为 "ref"，然后再把其中的 label 改为 "mygo"
    content = ""
    with open(f"{name}/{ref}/mygocmd", "r") as f:
        content = f.read()
        print(f"Change {ref} to ref, {label} to mygo in {name}/{ref}/mygocmd")
        content = content.replace(ref, "ref").replace("_base." + label, ".mygo")
    with open(f"{name}/{ref}/mygocmd", "w") as f:
        f.write(content)

    # 重命名 ref 文件夹为 "ref" 文件夹
    print(f"Rename {name}/{ref} to {name}/ref")
    os.rename(f"{name}/{ref}", f"{name}/ref")
